import request from '@/utils/request'

// ***********************Task

export function taskQueryList(params) {
  return request({
    url: '/task/queries',
    method: 'get',
    params // 通过url传参
  })
}

export function taskDelete(params) {
  return request({
    url: '/task/delete',
    method: 'post',
    data: params // json or data
  })
}

// ***********************Dataset

export function datasetQueryList(params) {
  return request({
    url: '/dataset/queries',
    method: 'get',
    params // 通过url传参
  })
}

export function datasetDelete(params) {
  return request({
    url: '/dataset/delete',
    method: 'post',
    data: params // json or data
  })
}

// ***********************Experiment

export function expQueryList(params) {
  return request({
    url: '/experiment/queries',
    method: 'get',
    params // 通过url传参
  })
}

export function expDelete(params) {
  return request({
    url: '/experiment/delete',
    method: 'post',
    data: params // json or data
  })
}



// ********************* Job
export function jobInsert(params) {
  return request({
    url: '/job/insert',
    method: 'post',
    data: params // json or data
  })
}

export function jobQueryOne(params) {
  return request({
    url: '/job/query',
    method: 'get',
    params
  })
}


export function jobQueryList(params) {
  return request({
    url: '/job/queries',
    method: 'get',
    params // 通过url传参
  })
}


export function jobDelete(params) {
  return request({
    url: '/job/delete',
    method: 'post',
    data: params // json or data
  })
}

