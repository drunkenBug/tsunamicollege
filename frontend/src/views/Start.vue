
<style media="screen">

</style>

<template>
    <div class="face-image">

    </div>
    <div class="face">

    </div>

    <h1 >Tsunami College</h1>
    <h2>Zeig was du kannst</h2>

    <TopicPicker/>


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

    insertSkill(name,level){
      this.skillMap[name]=level
      console.log(this.skillMap);
      console.log(JSON.stringify(this.skillMap));
      var raw=JSON.stringify(this.skillMap)
      console.log("raw:",raw);
      var newcookie="skillmap="+raw
      document.cookie=newcookie
      localStorage[name]=level
    },
    welcome(){

      for (var item in localStorage){
        console.log(item);
      }

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
      console.log("welcome found skills:",this.skillMap);
      this.insertSkill("Plus",90)


    },

  }
}
</script>
