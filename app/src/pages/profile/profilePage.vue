<template>
    <q-page class="column items-center">
        <q-img to="/schedule" src="/src/assets/ProfilePic.webp"
        height="100px"
        width="100px"
        style="border-radius: 50%; margin-top: 10px;"></q-img>
                <q-btn icon="camera_alt"
                 to="/profile/replace-image"
                 color="primary"
                  style="
                  margin-left: 65px;
                  bottom: 25px;"
                  size="xs"
                  round
                  ></q-btn>
        <div style="font-weight: 500; font-size: 25px;" >{{ userInfo["name"] }}</div>
        <div>{{ userInfo["phone-number"] }}</div>
        <div>{{ userInfo["email"] }}</div>
        <div>{{ userInfo["birthday"] }}</div>
        <q-list class="custom-list" >
          <div style="font-weight: 500; font-size: 20px;">Teams</div>
      <q-item v-for="team in userInfo['teams']" :key="team.name" class="list-item" :to="'/teaminfo/' + team.name">
          <q-img 
          :src="team.logo"
           width="45px"
            height="45px"
            class="rounded-borders"
            style="margin-right: 10px;"
            />
          <q-item-label style="font-size: 15px;">{{ team.name }}</q-item-label>
      </q-item>
    </q-list>
    <q-list class="custom-list" >
          <div style="font-weight: 500; font-size: 20px;">Leagues</div>
      <q-item v-for="league in userInfo['leagues']" :key="league" class="list-item" :to="'/leagueinfo/' + league">
          <q-img 
          :src="'/src/assets/' + league + '.jpeg'"
           width="45px"
            height="45px"
            class="rounded-borders"
            style="margin-right: 10px;"
            />
          <q-item-label style="font-size: 15px;">{{ league }}-League</q-item-label>
      </q-item>
    </q-list>     
    <q-item class="button-container">
      <q-btn unelevated 
      color="primary"
       label="Sign Out"
        to="/" 
        class="q-pa-sm q-ma-sm rounded-edges"
        style="width: 70vw;" />
      <q-btn unelevated color="primary" 
      label="Edit Profile"
      to="/profile/edit-profile"
       class="q-pa-sm q-ma-sm rounded-edges" 
       style="width: 70vw;"/>
    </q-item>
  </q-page>
</template>

<script setup lang="ts">
import {useProfileStore} from 'app/src/stores/profileStore';
import { ref } from 'vue';

const profileStore = useProfileStore();

const userInfo = ref({
  'name': profileStore.profile.name,
  'phone-number': profileStore.profile.number,
  'email': profileStore.profile.email,
  'birthday': profileStore.profile.birthday,
  'teams': profileStore.profile.teams, 
  'leagues': profileStore.profile.leagues
});

</script>

<style>
.custom-list {
  width: 70vw;
  margin-bottom: 30px;
}
.list-item{
  align-items: center;
  border-top: 1px solid #ddd;
  border-bottom: 1px solid #ddd;
  margin-top: -1px; 
}
.button-container{
  margin-top: 10px;
  width: 80vw;
  justify-content: space-around;
  align-items: center;
}
</style>