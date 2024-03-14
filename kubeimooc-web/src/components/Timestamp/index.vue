<template>
  <span>
    {{ dateTimeFormat(data) }}
  </span>
</template>

<script>
export default {
  props: {
    data: {
      type: Number
    }
  },
  name: "Timestamp",
  methods: {
    dateTimeFormat(createTime) {
      let createTimestampStr = createTime.toString()
      if (createTimestampStr.length <= 10) {
        createTimestampStr = createTimestampStr + "000"
      }
      let createTimestampInt = new Number(createTimestampStr)
      return this.getDateDiff(Date.parse(new Date(createTimestampInt)))
    },
    /**
     * @param {number} dateTimeStamp
     * @returns {string}
     */
    getDateDiff(dateTimeStamp) {
      const minute = 1000 * 60
      const hour = minute * 60
      const day = hour * 24
      const month = day * 30
      const now = new Date().getTime()
      const diffValue = now - dateTimeStamp
      if (diffValue < 0) {
        return
      }
      const monthC = diffValue / month
      const weekC = diffValue / (7 * day)
      const dayC = diffValue / day
      const hourC = diffValue / hour
      const minC = diffValue / minute
      let result = ''
      if (monthC >= 1) {
        result = '' + parseInt(1) + '月前'
      } else if (weekC >= 1) {
        result = '' + parseInt(weekC) + '周前'
      } else if (dayC >= 1) {
        result = '' + parseInt(dayC) + '天前'
      } else if (hourC >= 1) {
        result = '' + parseInt(hourC) + '小时前'
      } else if (minC >= 1) {
        result = '' + parseInt(minC) + '分钟前'
      } else {
        result = '刚刚'
      }
      return result
    }
  }
}
</script>

<style scoped>

</style>
