import MeCab
import sys
import gensim

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
		node = mecab.barseToNode(title)
		node = node.next
		while node:
			if node.next == None:
				break

			# 形態素の抽出
			text = node.feature.split(",")[6]


if __name__ == "__main__":
	main()