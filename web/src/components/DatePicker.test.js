import DatePicker from "./DatePicker.vue";
import { shallowMount, createLocalVue } from "@vue/test-utils";
import Vuetify from "vuetify";

describe("@components/date-picker", () => {
  it("exports a valid component", () => {
    const vuetify = new Vuetify();
    const localVue = createLocalVue();

    const wrapper = shallowMount(DatePicker, localVue, vuetify);

    expect(wrapper.contains("v-date-picker-stub")).toBe(true);
  });
});
