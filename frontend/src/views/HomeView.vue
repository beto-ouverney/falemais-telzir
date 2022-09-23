<style lang="scss" scopped>
@import "./styles/HomeView";
</style>
<template>
  <div>
    <div class="home">
      <div class="top-home">
        <img alt="Gopher Image" src="../assets/govue.webp" />
        <h1>FaleMais Telzir</h1>
      </div>
    </div>
    <FiltersField />
    <TablePlansCost :show="showTable" />
    <LoadingComponent :show="isLoading" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import FiltersField from "@/components/home/FiltersField.vue";
import TablePlansCost from "@/components/home/TablePlansCost.vue";
import { mapState } from "vuex";
import LoadingComponent from "@/components/shared/LoadingComponent.vue";

export default defineComponent({
  name: "HomeView",
  components: {
    LoadingComponent,
    FiltersField,
    TablePlansCost,
  },
  computed: {
    ...mapState(["isLoading", "showTable"]),
  },
});
</script>
