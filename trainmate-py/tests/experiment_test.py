import sys
import pytest
# from utils import distances
import requests

job_url = "http://localhost:8000/api/v1/experiment/"


def test_create():
    # jaccard_coefficient = distances.jaccrad(a, b)
    resp = requests.post(url=job_url + "insert", json={
        "name": "siamese_v20211121",
        "dataset_id": "1212121"
    })
    print(resp.status_code)
    print(resp.json())


def test_query():
    resp = requests.get(url=job_url + "query", json={
    })
    print(resp.status_code)
    print(resp.json())