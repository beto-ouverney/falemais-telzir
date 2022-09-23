<template>
  <div>
    <button type="button" @click="$emit('clickBtn')">
      {{ labelBtn }}
    </button>
  </div>
</template>

<script>
export default {
  name: "BtnCalc",
  props: {
    labelBtn: String,
  },
};
</script>
