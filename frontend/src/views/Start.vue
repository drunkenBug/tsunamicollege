<template>
  <div class="home">
    <div class="motto">
      <div class="face">
        <h1 >Tsunami College</h1>

      </div>

      <h2>Zeig was du kannst</h2>
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
