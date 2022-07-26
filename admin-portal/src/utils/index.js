/**
 * Created by PanJiaChen on 16/11/18.
 */

/**
 * Parse the time to string
 * @param {(Object|string|number)} time
 * @param {string} cFormat
 * @returns {string | null}
 */
export function parseTime(time, cFormat) {
  if (arguments.length === 0 || !time) {
    return null
  }
  const format = cFormat || '{y}-{m}-{d} {h}:{i}:{s}'
  let date
  if (typeof time === 'object') {
    date = time
  } else {
    if ((typeof time === 'string')) {
      if ((/^[0-9]+$/.test(time))) {
        // support "1548221490638"
        time = parseInt(time)
      } else {
        // support safari
        // https://stackoverflow.com/questions/4310953/invalid-date-in-safari
        time = time.replace(new RegExp(/-/gm), '/')
      }
    }

    if ((typeof time === 'number') && (time.toString().length === 10)) {
      time = time * 1000
    }
    date = new Date(time)
  }
  const formatObj = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay()
  }
  const time_str = format.replace(/{([ymdhisa])+}/g, (result, key) => {
    const value = formatObj[key]
    // Note: getDay() returns 0 on Sunday
    if (key === 'a') { return ['日', '一', '二', '三', '四', '五', '六'][value] }
    return value.toString().padStart(2, '0')
  })
  return time_str
}

/**
 * @param {number} time
 * @param {string} option
 * @returns {string}
 */
export function formatTime(time, option) {
  if (('' + time).length === 10) {
    time = parseInt(time) * 1000
  } else {
    time = +time
  }
  const d = new Date(time)
  const now = Date.now()

  const diff = (now - d) / 1000

  if (diff < 30) {
    return '刚刚'
  } else if (diff < 3600) {
    // less 1 hour
    return Math.ceil(diff / 60) + '分钟前'
  } else if (diff < 3600 * 24) {
    return Math.ceil(diff / 3600) + '小时前'
  } else if (diff < 3600 * 24 * 2) {
    return '1天前'
  }
  if (option) {
    return parseTime(time, option)
  } else {
    return (
      d.getMonth() +
      1 +
      '月' +
      d.getDate() +
      '日' +
      d.getHours() +
      '时' +
      d.getMinutes() +
      '分'
    )
  }
}

/**
 * @param {string} url
 * @returns {Object}
 */
export function param2Obj(url) {
  const search = decodeURIComponent(url.split('?')[1]).replace(/\+/g, ' ')
  if (!search) {
    return {}
  }
  const obj = {}
  const searchArr = search.split('&')
  searchArr.forEach(v => {
    const index = v.indexOf('=')
    if (index !== -1) {
      const name = v.substring(0, index)
      const val = v.substring(index + 1, v.length)
      obj[name] = val
    }
  })
  return obj
}
export function formatSize(size) {
  if (size.indexOf('m') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 0.001
    return size
  } else if (size.indexOf('Ki') !== -1) {
    size = size.substring(0, size.length - 2)
    size = size * 1024
    return size
  } else if (size.indexOf('Mi') !== -1) {
    size = size.substring(0, size.length - 2)
    size = size * 1024 * 1024
    return size
  } else if (size.indexOf('Gi') !== -1) {
    size = size.substring(0, size.length - 2)
    size = size * 1024 * 1024 * 1024
    return size
  } else if (size.indexOf('Ti') !== -1) {
    size = size.substring(0, size.length - 2)
    size = size * 1024 * 1024 * 1024 * 1024
    return size
  }
  if (size.indexOf('k') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000
    return size
  }
  if (size.indexOf('m') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000 * 1000
    return size
  }

  if (size.indexOf('g') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000 * 1000 * 1000
    return size
  }

  if (size.indexOf('t') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000 * 1000 * 1000 * 1000
    return size
  }

  if (size.indexOf('p') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000 * 1000 * 1000 * 1000 * 1000
    return size
  }
  if (size.indexOf('e') !== -1) {
    size = size.substring(0, size.length - 1)
    size = size * 1000 * 1000 * 1000 * 1000 * 1000 * 1000
    return size
  }
  else {
    return size
  }
}
export function formatDuring(mss) {
  mss = mss * 1000
  var days = parseInt(mss / (1000 * 60 * 60 * 24))
  days = days === 0 ? '' : days + 'd'
  var hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  hours = hours === 0 ? '' : hours + 'h'
  var minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60))
  minutes = minutes === 0 ? '' : minutes + 'm'
  var seconds = Math.round((mss % (1000 * 60)) / 1000) + 's'
  seconds = seconds === 0 ? '' : seconds
  return days + hours + minutes + seconds
}

