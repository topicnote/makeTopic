import MeCab
import gensim
import numpy
import os


class TopicCorpus():
	def __init__(self):
		self.modelPath = os.environ['NLP_MODEL_PATH']
		self.wordModelPath = self.modelPath+'/ja.bin'
		self.topicModelPath = self.modelPath+'/topic.bin'
		# 単語モデル、トピックモデル（トピック空間）の読み込み
		self.wordModel = gensim.models.Word2Vec.load(self.wordModelPath)
		self.topicModel = gensim.models.Word2Vec.load(self.topicModelPath)
		# MeCabをセット
		mecabPath = os.environ['MECAB_DIC_PATH']
		self.mecab = MeCab.Tagger("-d "+mecabPath)
		# topicのしきい値を設定
		self.threshold = 0.9
	

	def getNewsVector(self, newsTitle):
		topicVector = numpy.zeros(300)
		node = self.mecab.parseToNode(newsTitle)
		node = node.next
		while node:
			if node.next == None:
				break
			# 単語のVector化と重み付けをしてtopicVectorに加算
			word = node.feature.split(",")[6]
			score = self.wordScore(node)
			try:
				wordVector = self.wordModel[word]*score
			except :
				wordVector = numpy.zeros(300)				
			topicVector = topicVector + wordVector
			node = node.next
		return topicVector

	# 既存のtopicVectorに追加されたVectorの要素を追加して更新
	def updateTopicVector(self, newsVector, TopicID):
		# self.topicModel[TopicID] = numpy.mean(self.topicModel[TopicID] + newsVector)
		pass


	def addNewTopic(self, newsVector):
		newTopicID = len(self.topicModel.wv.vocab)-7
		# newTopicID = 0
		self.topicModel.wv.add(str(newTopicID), newsVector)
		return newTopicID


	def getTopicID(self, newsTitle):
		newsVector = self.getNewsVector(newsTitle)
		nearestTopic = self.topicModel.most_similar([newsVector],[],1)
		# print(nearestTopic)
		# nearestTopic:[(string)TopicID, (float?)distance]
		if abs(nearestTopic[0][1]) > self.threshold:
			self.updateTopicVector(newsVector, nearestTopic[0][1])
			return nearestTopic[0][0]
		else:
			newTopicID = self.addNewTopic(newsVector)
			return newTopicID
		
	def wordScore(self, node):
		# if node.feature.split(",")[1] == "固有名詞":
		# 	score = 2
		# elif node.feature.split(",")[1] in {"句点","格助詞"}: #適宜条件追加
		# 	score = 0
		# else:
		# 	score = 1
		# return score
		return 1

if __name__ == "__main__":
	tc = TopicCorpus()
	file = open(tc.modelPath+"/newsList.txt")
	newsList = file.readlines()
	newsTitle = None

	for newsTitle in newsList:
		topicID = tc.getTopicID(newsTitle)
		print(topicID)

	tc.topicModel.save(tc.topicModelPath)