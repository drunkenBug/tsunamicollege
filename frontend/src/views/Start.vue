<template>
  <div class="home">
    <div class="motto">
      <div class="face">
        <h1 >Tsunami College</h1>

      </div>

      <p>Zeig was du kannst</p>
    </div>
    <TopicPicker/>


  </div>
</template>

<script>
import TopicPicker from '../components/TopicPicker'
// import VueCookies from "vue-cookies"

export default {
  name: 'Start',
  components: {
    TopicPicker,
    // VueCookies,
  }
}
</script>
