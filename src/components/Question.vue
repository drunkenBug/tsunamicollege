<template>
  <div class="block">
    <div class="question">
      <h2>{{content.title}}</h2>
      <input type="text" v-model='input' ref="input" @keyup.enter='onEnter'>
      <button type="button" name="button" @click='go' >
        <div class="go">
          GO
        </div>
      </button>
      <p>{{result}}</p>
    </div>

  </div>
</template>

<script>
  export default{
    props:[
      'content',
    ],
    data(){
      return {
        result:"...",
        input:""
      }
    },
    methods:{
      go(){
        if (this.input==""){
          return
        }
        if (this.input==this.content.solution){
          this.result=this.content.solution+" is correct!"
          this.input=""
          this.$emit('solved')
        }else{
          this.result="false!"
        }
      },
      onEnter(){
        console.log("enter")
        this.go()
      },
      start(){
        this.$refs.input.focus()
        this.result="..."
      }
    },
    created(){
      console.log("created")
    }
  }
</script>

<style>
  .block{
    background: white;
    position:relative;
    left:50%;
    transform:translate(-50%,0);
    padding:2rem;
    border-radius:2rem;
  }


</style>
