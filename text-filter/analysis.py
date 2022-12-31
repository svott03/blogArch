from transformers import pipeline
import sys

sentiment_pipeline = pipeline("sentiment-analysis")

data = [' '.join(sys.argv[1:])]
# print(data)

print(sentiment_pipeline(data)[0]['label'])