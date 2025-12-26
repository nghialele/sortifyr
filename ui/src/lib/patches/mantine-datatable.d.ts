import { CollapseProps, MantineTheme } from "@mantine/core";
import { ReactNode } from "react";

// https://github.com/icflorescu/mantine-datatable/issues/651
declare module "mantine-datatable" {
  // https://icflorescu.github.io/mantine-datatable/type-definitions/
  // Expand the types as needed

  interface Column<T> extends Record<string, unknown> {
    accessor: keyof T | string;
    title?: string;
    width?: number | string;
    render?: (row: T) => ReactNode;
    sortable?: boolean;
    textAlign?: "left" | "right";
  }

  export interface DataTableProps<T> {
    idAccessor?: keyof T;
    columns: Column<T>[];
    records: T[];
    rowExpansion?: {
      content: (params: { record: T, index: number, collapse: () => void }) => ReactNode,
      collapseProps?: CollapseProps,
      allowMultiple?: boolean,
    };

    noHeader?: boolean;
    noRecordsText?: string;
    backgroundColor?: string;
    minHeight?: number;
    maxHeight?: number;
    height?: number;
    borderRadius?: "xs" | "sm" | "md" | "lg" | "xl";
    withTableBorder: boolean;
    highlightOnHover?: boolean;

    fetching?: boolean;
    customLoader?: ReactNode;
    onScrollToBottom?: () => void;

    styles: {
      root: (theme: MantineTheme) => Record<string, string>,
      header: (theme: MantineTheme) => Record<string, string>,
    };
  }

  export function DataTable<T>(props: DataTableProps<T>): React.ReactNode;
}
