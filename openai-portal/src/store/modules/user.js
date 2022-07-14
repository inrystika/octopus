import { login, logout, getInfo, getSpace } from '@/api/Home'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
import { Message } from 'element-ui'

const getDefaultState = () => {
  return {
    token: '',
    name: '',
    avatar: '',
    id: '',
    workspaces: [],
    progressId: undefined
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_ID: (state, id) => {
    state.id = id
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_SPACE: (state, workspaces) => {
    state.workspaces = workspaces
  },
  SET_PROGRESSID: (state, progressId) => {
    state.progressId = progressId
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { email, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: email.trim(), password: password }).then(response => {
        if (response.success) {
          const { token } = response.payload
          commit('SET_TOKEN', token)
          setToken(token)
          resolve('success')
        } else { resolve('failed') }
      }).catch(error => {
        Message({
          message: '未授权,请求没有操作的有效身份凭证',
          type: 'error',
          duration: 5 * 1000
        })
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(getToken()).then(response => {
        const { data } = response
        if (!data) {
          return reject('验证失败，请重新登录。')
        }
        const { fullName, id } = data.user
        commit('SET_NAME', fullName)
        commit('SET_ID', id)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  },
  getSpace({ commit, state }) {
    getSpace(state.id).then(response => {
      let data = []
      if (response.payload !== null) { data = response.payload.workspaces }
      data.forEach(item => {
        if (item.id === 'default-workspace') {
          item.name = '默认群组'
        }
      })
      commit('SET_SPACE', data)
      // var WORKSPACES = JSON.stringify(state.workspaces)
      // localStorage.setItem('WORKSPACES', WORKSPACES)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

