from transformers import pipeline
import sys

sentiment_pipeline = pipeline("sentiment-analysis")

data = []