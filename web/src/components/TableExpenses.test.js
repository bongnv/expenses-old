import TableExpenses from "./TableExpenses";
import { shallowMount } from "@vue/test-utils";

describe("@components/TableExpenses", () => {
  it("exports a valid component", () => {
    expect(shallowMount(TableExpenses).exists()).toBe(true);
  });
});
