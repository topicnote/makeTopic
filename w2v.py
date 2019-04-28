import MeCab
import sys
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


def main():
	if len(sys.argv) < 2:
		return -1
	# 単語モデル、トピックモデル（トピック空間）の読み込み
	wordModel = gensim.models.Word2Vec.load('ja.bin')
	topicModel = gensim.models.Word2Vec.load('topic.bin')

	# MeCabとニュースのタイトル群をセット
	mecab = MeCab.Tagger("-d $MECAB_DIC_PATH")
	file = open("./title.txt")
	lines = file.readlines()

	# 各行(タイトル)について単語で分割、Vector化、種類ごとに重み付けして総和を取る
	for title in lines:
		topicVector = numpy.zeros()
		node = mecab.barseToNode(title)
		node = node.next
		while node:
			if node.next == None:
				break

			# 単語のVector化と重み付けをしてtopicVectorに加算
			word = node.feature.split(",")[6]
			score = wordScore(node)
			wordVector = wordModel[word]*score

			topicVector = topicVector + wordVector

			node = node.next

		# 最近傍トピックの取得(topicID, distance)
		nearestTopic = topicModel.most_similar([topicVector],[],1)
# 最近傍トピックとの距離がしきい値以下ならそれに含み、そうでないなら新しくIDを振る
# 新トピックを生成した場合topicModelにそれを追加（やり方は要調査）
# すべてのニュースにトピックIDをつけたらIDを順番に標準出力
# topicModelを保存して終了
if __name__ == "__main__":
	main()