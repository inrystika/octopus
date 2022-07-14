const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  workspaces: state => state.user.workspaces,
  id: state => state.user.id,
  progressId: state => state.user.progressId

}
export default getters
