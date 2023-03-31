

ml-mate service helper trainer record experiment, dataset and performance

trainmate-sdk(ml-mate-py)
trainmate-fe(ml-mate-frontend)
trainmate-server(ml-mate-backend)

<!-- performance： task
- perf-id
- job-id
- metrics  # select job_id on testset -->

experiment:  task + model
<!-- - parent-id : perf-id -->
- exp-id
- exp-name  # prefix dataset_name+model_name
- dataset-id
- desp

job:  model+hyparamter+dataset
- parent-id: exp-id
- job-id
- config(hyperparameter)
- model_url
- history  # 用于画图
- metrics  # on trainset and valset
- metric avg， std .....


dataset:
- parent-id: task-id
- dataset-id
- desp # 描述
- name  # prefix=task_name
- version  # 版本
- source
- store_type

task:
- task-id
- desp
- name



### apis
##### job
- create 一个exp中，job的name不能重复，如果重复则create失败
- update 根据job-id/name+exp-id来更新
- delete  by job-id/name+exp-id
- query  by job-id/name+exp-id


## todo
# sqlite
# config  from yml


Reference
- [1]git clone https://github.com/piexlmax/gin-vue-admin.git