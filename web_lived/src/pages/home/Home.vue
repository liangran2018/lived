<template>
    <div id="home">
      <home-time
        :data="data"
        :time="time"
        :user="user"
      ></home-time>
    </div>
</template>

<script>
import HomeTime from './components/Time'
import axios from 'axios'
export default {
  name: 'Home',
  components: {
    HomeTime
  },
  data () {
    return {
      data: {
        type: Object
      },
      user: {
        type: Object
      },
      time: {
        type: Object
      }
    }
  },
  methods: {
    getHomeInfo () {
      axios.get('/api/newgame')
        .then(this.getHomeInfoSucc)
    },
    getHomeInfoSucc (res) {
      let data = res.data.data
      this.data = data
      this.user = data.user
      this.time = data.time
      let hours = this.time.hour
      let minutes = this.time.minute
      if (hours < 10) {
        this.time.hour = '0' + hours
      } else {
        this.time.hour = hours
      }
      if (minutes < 10) {
        this.time.minute = '0' + minutes
      } else {
        this.time.minute = minutes
      }
    }
  },
  mounted () {
    this.getHomeInfo()
  }
}
</script>

<style lang="stylus" scoped>

</style>
