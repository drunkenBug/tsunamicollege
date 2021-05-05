<template>
  <div class="about">
    <h1>Was ist Tsunamicollege</h1>
    <h3>Tsunamicollege versucht Schule besser zu machen.</h3>
    <div class="info">

    <h3>
      <p>damit wir Fragen stellen können, die den richtigen schwierigkeitsgrad haben, speichern wir
      wie gut du in den einzelnden Themen bist in deinem Browser. Diese Daten werden von uns nicht gesammelt.</p>
      <p>wenn du ganz von vorne anfangen möchtest kannst du alle gespeicherten Daten löschen:</p>
    </h3>
  </div>

    <!-- <div class="entrance">
      <h3>Du kannst Tsunamicollege besser machen,<p></p>
         indem du eine neue Frage zu unserer Sammlung hinzufügst:</h3>
      <button type="button" name="button">
        <router-link to="/Editor">Neue Testfrage hinzufügen</router-link>
      </button>
    </div> -->
    <button type="button" name="button" @click="emptyStorage"><h3>lösche alle daten</h3> </button>

  </div>
</template>

<style media="screen">
  button{
    background: transparent;
  }
  .info{
    margin:auto;
    width:60%;
  }
</style>
<script>
  export default{
  methods:{
    emptyStorage(){
      var topics=["Plus","Minus","Ein mal eins","Geteilt"]
      for (var i in topics){
        localStorage[topics[i]]=0
      }
      alert("daten wurden gelöscht")
    }
  }
}

</script>
