import { TextInput, TextInputProps } from "@mantine/core"
import { LuSearch } from "react-icons/lu"

type Props = TextInputProps

export const Search = (props: Props) => {
  return (
    <TextInput
      leftSection={<LuSearch />}
      radius="lg"
      styles={{
        input: {
          background: "var(--mantine-color-background-0)",
          borderColor: "var(--mantine-color-background-0)",
        }
      }}
      {...props}
    />
  )
}
