import request from '@/utils/request'

export function getTypeList() {
  return request({
    url: '/admin/getTypeList',
    method: 'GET'
  })
}
export function delType(data) {
  return request({
    url: '/admin/DelType',
    method: 'POST',
    data
  })
}

export function addType(data) {
  return request({
    url: '/admin/AddType',
    method: 'POST',
    data
  })
}

export function editType(data) {
  return request({
    url: '/admin/EditType',
    method: 'POST',
    data
  })
}

export function getUrlList(data) {
  return request({
    url: '/admin/getUrlList',
    method: 'POST',
    data
  })
}

export function delUrl(data) {
  return request({
    url: '/admin/DelUrl',
    method: 'POST',
    data
  })
}

export function addUrl(data) {
  return request({
    url: '/admin/AddUrl',
    method: 'POST',
    data
  })
}

export function editUrl(data) {
  return request({
    url: '/admin/EditUrl',
    method: 'POST',
    data
  })
}
