import { cn } from "@/lib/utils";
import { Stack, StackProps, Title } from "@mantine/core";

type PageProps = StackProps

export function Page({ className = "", ...props }: PageProps) {
  return (
    <Stack p="lg" className={cn("h-full", className)} {...props} />
  )
}

type PageTitleProps = {
  title: string;
  description?: string;
} & StackProps

export const PageTitle = ({ title, description, ...props }: PageTitleProps) => {
  return (
    <Stack py="md" gap={0} h="92px" {...props}>
      <Title order={1} size="h2" >{title}</Title>
      <p className="text-muted">{description}</p>
    </Stack>
  )
}


type SectionProps = StackProps

export const Section = ({ className, ...props }: SectionProps) => {
  return <Stack p="md" gap="xs" className={cn("flex-1 bg-white rounded-xl md:overflow-hidden", className)} {...props} />
}

type SectionTitleProps = {
  title: string;
  description?: string;
} & StackProps

export const SectionTitle = ({ title, description, ...props }: SectionTitleProps) => {
  return (
    <Stack gap={0} {...props}>
      <Title order={2} size="h3">{title}</Title>
      <p className="text-muted whitespace-pre-wrap">{description}</p>
    </Stack>
  )
}
