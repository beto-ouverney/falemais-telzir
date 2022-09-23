<template>
  <div>
    <label
      >{{ label }}
      <input v-model="value" type="number" @input="handleChange" />
    </label>
  </div>
</template>
<script>
export default {
  name: "InputNumber",
  props: {
    label: String,
  },
  data() {
    return {
      value: 0,
    };
  },
  methods: {
    handleChange() {
      console.log(this.value);
      this.$emit("updateValue", this.value);
    },
  },
};
</script>
