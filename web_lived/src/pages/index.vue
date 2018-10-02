<template>
    <div id="index">
      <img class="background_img" src="../../static/index.jpg"/>
      <div class="fiexdbox">
        <div class="name">L I V E D</div>
        <router-link to="/home">
          <div class="button" style="margin-right: .6rem">新游戏</div>
        </router-link>
        <div class="button">载入</div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'index'
}
</script>

<style lang="stylus" scoped>
.background_img
  width:100%;
.fiexdbox
  position: fixed;
  font-size: 0;
  margin: auto;
  left: 0;
  right: 0;
  top: 6rem;
  bottom: 0;
  width: 5.84rem;
  height: 12.3rem;
  .name
    font-size: 1.2rem;
    color: #662603;
    margin-bottom: 4rem;
    text-align: center;
  .button
    display inline-block;
    width:2.6rem;
    height: 1rem;
    line-height: 1rem;
    text-align: center;
    font-size: .6rem;
    color: #ffffff;
    border-radius: .2rem;
    border: .01rem solid #95a7b3;
    background: #1b1b1bc9;
    font-family: "新宋体";

</style>
