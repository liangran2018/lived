<template>
    <div id="home">
      <home-time :data="data"></home-time>
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
      data: Object
    }
  },
  methods: {
    getHomeInfo () {
      axios.get('/api/newgame')
        .then(this.getHomeInfoSucc)
    },
    getHomeInfoSucc (res) {
      var data = res.data.data
      this.data = data
    }
  },
  mounted () {
    this.getHomeInfo()
  }
}
</script>

<style lang="stylus" scoped>

</style>
