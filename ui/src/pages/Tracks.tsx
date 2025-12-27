import { Page, PageTitle, Section } from "@/components/atoms/Page";
import { Segment } from "@/components/molecules/Segment";
import { TrackAdded } from "@/components/track/TrackAdded";
import { TrackDeleted } from "@/components/track/TrackDeleted";
import { TrackHistory } from "@/components/track/TrackHistory";
import { Center, Group } from "@mantine/core";
import { ReactNode, useMemo, useState } from "react";
import { LuSquareStack, LuTextSearch } from "react-icons/lu";

type ViewOption = "history" | "added" | "deleted"
type View = { value: ViewOption, label: string, icon: ReactNode, }

const views: View[] = [
  {
    value: "history",
    label: "History",
    icon: <LuTextSearch />,
  },
  {
    value: "added",
    label: "Added",
    icon: <LuSquareStack />,
  },
  {
    value: "deleted",
    label: "Deleted",
    icon: <LuSquareStack />,
  },
]

const renderView = (view: ViewOption) => {
  switch (view) {
    case "history": return <TrackHistory />
    case "added": return <TrackAdded />
    case "deleted": return <TrackDeleted />
  }
}

const storageKey = "sortifyr-track-view"

export const Tracks = () => {
  const [view, setView] = useState<ViewOption>(localStorage.getItem(storageKey) as ViewOption ?? "history")

  const handleSegment = (view: ViewOption) => {
    localStorage.setItem(storageKey, view)
    setView(view)
  }

  const renderedView = useMemo(() => renderView(view), [view])

  return (
    <Page>
      <Group justify="space-between">
        <PageTitle
          title="Tracks"
          description="All track related tools."
        />
        <Segment
          data={views.map(v => ({
            value: v.value,
            label: (
              <Center style={{ gap: 4 }}>
                {v.icon}
                <p>{v.label}</p>
              </Center>
            )
          }))}
          value={view}
          onChange={e => handleSegment(e as ViewOption)}
        />
      </Group>

      <Section>
        {renderedView}
      </Section>

    </Page>
  )
}
