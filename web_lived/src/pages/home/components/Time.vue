<template>
  <div>
    <img class="user_img" src="../../../../static/user_img.png"/>
    <ul class="time_list">
      <li class="time_item hour">{{this.hour}}:{{this.minute}}</li>
      <li class="time_item temprature">{{this.data.temprature}}</li>
      <li class="time_item overday">{{this.data.time.overday}}</li>
      <li class="time_item season">{{this.data.season}}</li>
      <li class="time_item weather">{{this.data.weather}}</li>
    </ul>
    <ul class="user_list">
      <li class="user_item hurt">{{this.data.user.Hurt}}</li>
      <li class="user_item ill">{{this.data.user.Ill}}</li>
      <li class="user_item thirst">{{this.data.user.Thirst}}</li>
      <li class="user_item hungry">{{this.data.user.Hungry}}</li>
      <li class="user_item mood">{{this.data.user.Mood}}</li>
      <li class="user_item blood">{{this.data.user.Blood}}</li>
      <li class="user_item wake">{{this.data.user.Wake}}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Time',
  props: {
    data: Object
  },
  data () {
    return {
      hour: '',
      minute: ''
    }
  },
  mounted () {
    let hours = this.data.time.hour
    let minutes = this.data.time.minute
    if (hours < 10) {
      hours.toString()
      this.hour = '0' + hours
    } else {
      this.hour = hours.toString()
    }
    if (minutes < 10) {
      minutes.toString()
      this.minute = '0' + minutes
    } else {
      this.minute = minutes.toString()
    }
  }
}
</script>

<style lang="stylus" scoped>
.user_img
  position absolute
  width: 1.5rem;
  border: .02rem solid #231a11;
  border-radius: .2rem;
  background-color: #86603b91;
.time_list
  background-color: #86603b91;
  height: .6rem;
  margin-left: 1.5rem;
  color: #fff;
  .time_item
    float left
    height .6rem
    line-height .6rem
    background-position-x .05rem
    background-position-y 50%
    background-repeat no-repeat
    background-size .3rem .3rem
    text-indent .36rem
  .hour
    background-image url("../../../../static/time.png")
  .temprature
    background-image url("../../../../static/temperature.png")
  .overday
    background-image url("../../../../static/overday.png")
  .season
    background-image url("../../../../static/weather.png")
  .weather
    background-image url("../../../../static/weather.png")
.user_list
  background-color: #86603b91;
  height: .9rem;
  margin-left: 1.5rem;
  color: #fff;
  .user_item
    float left
    width .8rem
    height .9rem
    line-height .9rem
    background-position-x .05rem
    background-position-y 50%
    background-repeat no-repeat
    background-size .7rem .7rem
    text-align center
    border: .01rem solid #231a11;
    border-left: none;
  .hurt
    background-image url("../../../../static/hurt.png")
  .ill
    background-image url("../../../../static/ill.png")
  .hurt
    background-image url("../../../../static/hurt.png")
  .thirst
    background-image url("../../../../static/thirst.png")
  .hungry
    background-image url("../../../../static/hungry.png")
  .mood
    background-image url("../../../../static/mood.png")
  .blood
    background-image url("../../../../static/blood.png")
  .wake
    background-image url("../../../../static/wake.png")
</style>
