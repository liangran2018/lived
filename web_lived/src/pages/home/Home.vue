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
      console.log(data)
      this.data = data
    }
  },
  mounted () {
    this.getHomeInfo()
  }
}
</script>

<style scoped>
#home{
  font-size: 30px;
  padding: 20px;
}
</style>
