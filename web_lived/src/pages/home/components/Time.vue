<template>
  <div>
    <div>
      <div>{{this.data.season}}</div>
      <div>{{this.data.weather}}</div>
      <div>{{this.data.time.day}}</div>
      <div>{{this.data.time.hour}}:{{this.data.time.minute}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Time',
  props: {
    data: Object
  }
}
</script>

<style scoped>

</style>
