import MeCab
import sys
import gensim

def main(newsTitle):
	if len(sys.argv) < 2:
		return -1
	wordModel = gensim.models.Word2Vec.load('ja.bin')
	mecab = MeCab.Tagger("-d $MECAB_DIC_PATH")
	node = mecab.barseToNode(sys.argv[1])
	node = node.next
	while node:
		if node.next == None:
			break

		# 形態素の抽出
		text = node.feature.split(",")[6]


if __name__ == "__main__":
	main()