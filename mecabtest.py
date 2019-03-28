import MeCab
import sys

def main():
	if len(sys.argv) < 2:
		print("使い方: python3 mecabtest.py 解析したい文章")
		return 1
	mecab = MeCab.Tagger("-Ochasen")
	mecab.parseToNode('') #if not needed, remove.	
	node = mecab.parseToNode(sys.argv[1])
	while node:
		word = (node.surface).strip(node.next.surface)
		hinshi = node.feature.split(",")[0] #1,2,..により詳しい品詞や活用が格納されている
		print(word + " " + hinshi)
		node = node.next
	return 0

if __name__ == '__main__':
	main()