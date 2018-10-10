<template>
  <div>
    <img class="user_img" src="../../../../static/user_img.png"/>
    <ul class="time_list">
      <li class="time_item hour">{{this.time.hour}}:{{this.time.minute}}</li>
      <li class="time_item temprature">{{this.time.temprature}}</li>
      <li class="time_item overday">{{this.time.day}}</li>
      <li class="time_item season">{{this.data.season}}</li>
      <li class="time_item weather">{{this.data.weather}}</li>
      <li class="lv_exp">Lv.{{this.user.Lvl}}</li>
      <li class="lv_exp">Exp.{{this.user.Exp}}</li>
    </ul>
    <ul class="user_list">
      <li class="user_item hurt">{{this.user.Hurt}}</li>
      <li class="user_item ill">{{this.user.Ill}}</li>
      <li class="user_item thirst">{{this.user.Thirst}}</li>
      <li class="user_item hungry">{{this.user.Hungry}}</li>
      <li class="user_item mood">{{this.user.Mood}}</li>
      <li class="user_item blood">{{this.user.Blood}}</li>
      <li class="user_item wake">{{this.user.Wake}}</li>
    </ul>
    <img class="warehouse" src="../../../../static/warehouse.png" ref="warehouse"/>
    <img class="sleep1" src="../../../../static/sleep1.png"/>
    <img class="fire1" src="../../../../static/fire1.png" ref="fire1"/>
    <img class="medicine1" src="../../../../static/medicine1.png" ref="medicine1"/>
    <img class="tool1" src="../../../../static/tool1.png" ref="tool1"/>
    <img class="water1" src="../../../../static/water1.png" ref="water1"/>
  </div>
</template>

<script>
export default {
  name: 'Time',
  props: {
    data: Object,
    time: Object,
    user: Object
  },
  data () {
    return {
    }
  },
  mounted () {
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
  margin-left: 1.54rem;
  color: #fff;
  .lv_exp
    color: #000000;
    line-height: .6rem;
    margin-left: .2rem;
    float: left;
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
  margin-left: 1.54rem;
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
    margin-right: .03rem;
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
.sleep1
  width: 3rem;
  position: absolute;
  top: 5rem;
  right: .1rem;
.fire1
  width: 3rem;
  position: absolute;
  top: 5.5rem;
  left: .8rem;
  opacity: .6;
.medicine1
  width: 3.5rem;
  position: absolute;
  top: 7rem;
  right: .1rem;
  opacity: .6;
.water1
  width: 3rem;
  position: absolute;
  top: 8rem;
  left: 1rem;
  opacity: .6;
.tool1
  width: 3.5rem;
  position: absolute;
  top: 10rem;
  right: .1rem;
  opacity: .6;
.warehouse
  width: 3rem;
  position: absolute;
  top: 2.6rem;
  left: 0;
  opacity: .7;
</style>
