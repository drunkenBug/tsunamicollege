<template>
  <div class="test">
    <div class="face-image" style="z-index:-100;background:black">

    </div>
    <div ref="face" class="face-image" >
      <div id="star" ref='star'>
        <div class="color">

        </div>

      </div>

    </div>
    <div class="face">

    </div>

    <!-- <h1>{{topic}}</h1> -->
    <div v-if="questionsLoaded">
      <Question :content="question" ref="question" @failed="failed" @solved="solved" @next="next"/>
    </div>

  </div>


</template>

<script>
import Question from '../components/Question.vue'
import axios from 'axios'
export default {
  name: 'Test',
  components: {
    Question
  },
  data(){
    return {
      questionsLoaded:false,
      question:{},
      data:[],
      index:1,
      topic:this.$route.params.id,

      hist:[-1],

      diffscale:0.2,//how much does the difference in difficulty get accounted for
      reward:1,
      penalty:0.8,

    }
  },
  mounted(){
    if (localStorage[this.topic]==null){
      // console.log("no data on topic ",this.topic);
      localStorage[this.topic]=0
    }
    // console.log("skill level:",localStorage[this.topic]);
    var rating=Math.round(localStorage[this.topic])
    // console.log('hist: ',this.hist);
    this.getNewQuestion(this.topic,rating)
    //this.$refs.face.classList.add("right")
    // console.log(this.$refs.face);
  },
  methods:{
    joy(){
      this.$refs.star.classList.remove("impuls")
      setTimeout(()=>{
        this.$refs.star.classList.add('impuls')
      },10)
    },

    getTopicInfo(topic){
      var data={"op":"GETINFO","topic":topic,"no":0}
      axios({method:"POST",url:"/api/getTopicInfo",data:data,headers:{"content-type":"text/plain"}}).then(result=>{
        localStorage[topic+"Info"]=JSON.stringify(result.data);
      })
    },

    getNewQuestion(topic,rating){
      var data={'topic':topic,'rating':rating,'history':this.hist}
      axios({method:'POST',url:'api/new',data:data,headers:{'content-type':'text/plain'}}).then(result =>{
        this.question=result.data
        this.questionsLoaded=true

      }).catch(error=>{
        console.log(error);
      })
    },
    solved(){
      this.rating=parseFloat(localStorage[this.topic])
      var difference=this.question.rating-this.rating
      var change = Math.max(0,this.reward+difference*this.diffscale)
      this.rating+=change
      localStorage[this.topic]=this.rating
      this.joy()


    },
    failed(){
      this.rating=parseFloat(localStorage[this.topic])
      var difference=this.question.rating=this.rating
      var change = Math.max(0,this.penalty+difference*this.diffscale)
      this.rating-=change
      localStorage[this.topic]=this.rating

    },
    resetSkill(){
      localStorage[this.topic]=0
    },


    next(){
      this.hist.push(this.question.id)
      if (this.hist.length>10){
        this.hist=this.hist.slice(-10)
      }

      var rating=Math.round(localStorage[this.topic])
      // var hist = JSON.stringify(this.hist)
      this.getNewQuestion(this.topic,rating)
      this.$refs.question.start()
      // this.$refs.star.classList.remove("impuls")
    },

  }
}
</script>
