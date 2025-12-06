import { SegmentedControl, SegmentedControlProps } from "@mantine/core";

type Props = SegmentedControlProps

export const Segment = (props: Props) => {
  return (
    <SegmentedControl
      radius="lg"
      color="secondary.1"
      styles={{
        innerLabel: {
          color: "black",
        },
      }}
      {...props}
    />
  )
}
