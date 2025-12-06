import { Directory } from "@/lib/types/directory";
import { Side } from "@/lib/types/general";
import { Stack, StackProps } from "@mantine/core";
import { LoadingSpinner } from "../molecules/LoadingSpinner";
import { LinkTreeNode } from "./LinkTreeNode";
import { cn } from "@/lib/utils";

type Props = {
  roots: Directory[];
  side: Side;
  title: string;
  isLoading: boolean;
} & StackProps

export const LinkTree = ({ roots, side, title, isLoading, className, ...props }: Props) => {
  const child = () => {
    if (isLoading) return <LoadingSpinner />
    if (roots.length === 0) return null

    return (
      <>
        {roots.map(r => <LinkTreeNode key={r.id} directory={r} side={side} level={0} />)}
      </>
    )
  }

  return (
    <Stack p="md" pr={side === "left" ? "xl" : "md"} gap="md" justify="start" bg="background.0" className={cn("overflow-auto rounded-xl relative", className)} {...props}>
      <p className="text-muted">{title}</p>
      {child()}
    </Stack>
  )
}


