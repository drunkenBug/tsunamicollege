<template>
  <div class="element">
    <button :id ="getButtonId()"  class='topicButton'>
      <router-link :to="{name:'Test',params:{id:topicID}}" >
        <h3> {{topicID}}</h3>
      </router-link>
    </button>
  </div>



</template>

<script>
  export default{
    props:[
      'topicID',
    ],
    data(){
      return {
        input:""
      }
    },
    mounted(){
      console.log("created");
      console.log(localStorage[this.topicID])
      this.createStyle()
    },
    methods:{

      get_link(){
        var base="Test/"
        var tail="plusminus"
        var link=base+tail
        console.log(link);
        return "Test/"+this.title
      },
      getButtonId(){
        //generate a good id without whitespaces
        return (this.topicID+"button").replace(/\s+/g, '')
      },
      getTopicProgress(){
        //get the topic progress from 0 - 100
        return localStorage[this.topicID]*5
      },

      createStyle(){


        // create an new style
        const style = document.createElement('style');

        // append to DOM
        document.head.appendChild(style);


        // insert CSS Rule
        style.sheet.insertRule(`
            #`+this.getButtonId()+` {
              background: linear-gradient(90deg, var(--active) `+this.getTopicProgress()+`%, var(--passive) 0%);
            }
        `);
      }
    },
  }
</script>
