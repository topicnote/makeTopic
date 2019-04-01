#!/usr/bin/env python
# -*- coding:utf-8 -*-

import MeCab
import sys

def main():
	if len(sys.argv) < 2:
		print("使い方: python3 mecabtest.py 解析したい文章")
		return 1
	mecab = MeCab.Tagger("-Ochasen")
	node = mecab.parseToNode(sys.argv[1])
	index = 0
	gene = []
	while node:
		if node.next == None:
			break
		
		#解析対象の形態素を抽出(同じ文章を二回解析すると文字化けする謎現象発生中)
		text = (node.surface)
		nextText = (node.next.surface)
		text = text.replace(nextText, "")

		#for debug
		print(text)
		print(node.feature)

		#重み付け
		if node.feature.split(",")[1] == "固有名詞":
			score = 5
		elif node.feature.split(",")[1] in {"句点","格助詞"}: #適宜条件追加
			score = 0
		else:
			score = 1

		#配列に追加
		gene.append([text, score])
		
		#インクリメント
		node = node.next
		index = index + 1
	
	print(gene)
	return 0

if __name__ == '__main__':
	main()