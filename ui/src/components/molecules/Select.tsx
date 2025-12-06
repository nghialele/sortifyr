
import { Select as MSelect, SelectProps } from "@mantine/core"
import { LuSearch } from "react-icons/lu"

type Props = SelectProps

export const Select = (props: Props) => {
  return (
    <MSelect
      searchable
      allowDeselect
      leftSection={<LuSearch />}
      radius="lg"
      styles={{
        input: {
          background: "var(--mantine-color-background-0)",
          borderColor: "var(--mantine-color-background-0)",
        },
      }}
      className="grow"
      {...props}
    />
  )
}
