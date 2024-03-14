<template>
  <div>
    hello {{ message }}!
  </div>
</template>

<script>
export default {
  name: "ExampleIndex",

  data() {
    return {
      message: ""
    }
  },

  created() {
    this.$store.dispatch("example/examplePing").then(response => {
      let data = response.data
      this.message = data.message
    })
  }
}
</script>

<style scoped>

</style>
