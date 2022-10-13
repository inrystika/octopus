import Vue from 'vue'
const MyDirective = {}
export default MyDirective.install = function(vue, options) {
    Vue.directive('loadmore', {
        bind(el, binding) {
            const selectDom = el.querySelector('.el-select-dropdown .el-select-dropdown__wrap')
            selectDom.addEventListener('scroll', function() {
                    binding.value()
            })
        }
    })
}