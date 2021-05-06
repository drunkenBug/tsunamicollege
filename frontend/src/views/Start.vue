
<template>



    <div class="face-image">
      <div class="joy" id="star">
      </div>
    </div>
    <div class="face">


    </div>
    <h1>Tsunami College</h1>
    <h2>Zeig was du kannst</h2>

    <TopicPicker/>
    <div class="spacer">

    </div>

</template>
<script>
import TopicPicker from '../components/TopicPicker'

export default {
  name: 'Start',
  components: {
    TopicPicker,
    // VueCookies,
  },
  data(){
    return {
      skillMap:{}
    }
  },
  mounted(){
    this.welcome()
  },
  methods:{


    welcome(){



      var c=document.cookie;
      var cookies = c.split(";")
      var cookie =""
      for (cookie in cookies) {
        if (cookie.startsWith("skillmap=")){
          cookie=cookie.split("skillmap=")[0]
          cookie=cookie.split(";")[0]
          console.log(cookie)
          this.skillMap=JSON.parse(cookie)
          break
        }
      }
      // console.log("welcome found skills:",this.skillMap);


    },

  }
}
</script>
