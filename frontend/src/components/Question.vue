<template>
  <div class="bubble">
    <div class="question">

      <h2>{{content.title}}</h2>
      <input class="test" type="text" v-model='input' ref="input" @keyup.enter='onEnter'>
      <button class="test" type="button" name="button" @click='go' >
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
