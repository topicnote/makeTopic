import w2v
import numpy as np

def cos_sim(v1, v2):
    return np.dot(v1, v2) / (np.linalg.norm(v1) * np.linalg.norm(v2))


if __name__ == "__main__":
	tc = w2v.TopicCorpus()
	file = open(tc.modelPath+"/newsList.txt")
	newsList = file.readlines()
	newsTitle = None
	vec = None

	for i, newsTitle in enumerate(newsList):
		print(newsTitle)
		if i==0:
			vec = tc.getNewsVector(newsTitle)
			continue
		oldvec = vec
		vec = tc.getNewsVector(newsTitle)
		print(cos_sim(oldvec,vec))
