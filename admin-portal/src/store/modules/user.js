import { login } from '@/api/user'
import { getToken, setToken } from '@/utils/auth'
import { Message } from 'element-ui'
const getDefaultState = () => {
  return {
    token: getToken(),
    name: '',
    avatar: '',
    progress: 0
  }
}

const state = getDefaultState()

const mutations = {
  // RESET_STATE: (state) => {
  //   Object.assign(state, getDefaultState())
  // },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_PROGRESS: (state, progress) => {
    state.progress = progress
  },
  CLEAR_PROGRESS: (state, progress) => {
    state.progress = 0
  }
  // SET_NAME: (state, name) => {
  //   state.name = name
  // },
  // SET_AVATAR: (state, avatar) => {
  //   state.avatar = avatar
  // }
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
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

