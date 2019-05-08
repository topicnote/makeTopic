import MeCab
import gensim
import numpy


def wordScore(node):
	if node.feature.split(",")[1] == "固有名詞":
		score = 5
	elif node.feature.split(",")[1] in {"句点","格助詞"}: #適宜条件追加
		score = 0
	else:
		score = 1
	return score


class TopicCorpus():
	def __init__(self):
		# 単語モデル、トピックモデル（トピック空間）の読み込み
		self.wordModel = gensim.models.Word2Vec.load('/home/ubuntu/nlp/w2vmodel/ja.bin')
		self.topicModel = gensim.models.Word2Vec.load('/home/ubuntu/nlp/w2vmodel/topic.bin')
		# MeCabをセット
		self.mecab = MeCab.Tagger("-d /usr/lib/x86_64-linux-gnu/mecab/dic/mecab-ipadic-neologd")
		# topicのしきい値を設定
		self.threshold = 0.1
	

	def getNewsVector(self, newsTitle):
		topicVector = numpy.zeros(300)
		node = self.mecab.parseToNode(newsTitle)
		node = node.next
		while node:
			if node.next == None:
				break
			# 単語のVector化と重み付けをしてtopicVectorに加算
			word = node.feature.split(",")[6]
			score = wordScore(node)
			try:
				wordVector = self.wordModel[word]*score
			except :
				wordVector = numpy.zeros(300)				
			topicVector = topicVector + wordVector
			node = node.next
		return topicVector

	# 既存のtopicVectorに追加されたVectorの要素を追加して更新
	def updateTopicVector(self, newsVector, TopicID):
		self.topicModel[TopicID] = (self.topicModel[TopicID] + newsVector)/2


	def addNewTopic(self, newsVector):
		newTopicID = len(self.topicModel.vocab)
		self.topicModel.add(str(newTopicID), newsVector)
		return newTopicID


	def getTopicID(self, newsTitle):
		newsVector = self.getNewsVector(newsTitle)
		# nearestTopic:[(string)TopicID, (float?)distance]
		nearestTopic = self.topicModel.most_similar([newsVector],[],1)
		print(nearestTopic)
		# if nearestTopic[1] < self.threshold:
		# 	self.updateTopicVector(newsVector, nearestTopic[1])
		# 	return nearestTopic[0] + "*"
		# else:
		# 	newTopicID = self.addNewTopic(newsVector)
		# 	return str(newTopicID)
		newTopicID = self.addNewTopic(newsVector)
		return str(newTopicID)

if __name__ == "__main__":
	topicCorpus = TopicCorpus()
	file = open("/home/ubuntu/nlp/w2vmodel/newsList.txt")
	newsList = file.readlines()
	newsTitle = None

	for newsTitle in newsList:
		topicID = topicCorpus.getTopicID(newsTitle)
		print(topicID)

	topicCorpus.topicModel.save('/home/ubuntu/nlp/w2vmodel/topic.bin')