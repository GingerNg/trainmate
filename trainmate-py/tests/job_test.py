import sys
import pytest
# from utils import distances
import requests

# job_url = "http://localhost:8000/api/v1/job/"
job_url = "http://47.96.41.173:8666/api/v1/job/"


def test_create():
    # jaccard_coefficient = distances.jaccrad(a, b)
    resp = requests.post(url=job_url + "insert", json={
        "exp_id": "12121",
        "name": "siamese_v2021111121223",
        "metrics": {},
        "job_id": "",
        "config": {},
        "history": {}
        # "config": {"epoch": 20, "dropout": 0.1},
        # "metrics": {"acc": 0.99, "recall": 0.89}
    })
    print(resp.status_code)
    print(resp.json())


def test_query_by_id():
    # jaccard_coefficient = distances.jaccrad(a, b)
    resp = requests.get(url=job_url+"query", params={
        "job_id": "222ada85-0447-4908-bb96-55bc7bcf061d",
    })
    print(resp.status_code)
    print(resp.json())


def test_query_by_expid():
    # jaccard_coefficient = distances.jaccrad(a, b)
    # resp = requests.get(url=job_url+"querybyexpid", json={
    #     # "exp_id": "12121",
    # })
    resp = requests.get(url=job_url+"querybyexpid", params={
        "exp_id": "12121",
    })
    print(resp.status_code)
    print(resp.json())


def test_query_by_expid_and_name():
    # jaccard_coefficient = distances.jaccrad(a, b)
    resp = requests.get(url=job_url+"query", params={
        "exp_id": "12121",
        "name": "siamese_v20211121"
    })
    print(resp.status_code)
    print(resp.json())


def test_update():
    resp = requests.post(url=job_url+"update", json={
        "job_id": "222ada85-0447-4908-bb96-55bc7bcf061d",
        "name": "siamese_v20211121",
        "config": "121212"
    })
    print(resp.status_code)
    print(resp.json())


def test_delete():
    resp = requests.post(url=job_url+"delete", json={
        "job_id": "222ada85-0447-4908-bb96-55bc7bcf061d",
    })
    print(resp.status_code)
    print(resp.json())
