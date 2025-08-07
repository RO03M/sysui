import { useQuery } from "react-query";
import { api } from "./api";
import { Duration } from "../utils/duration";
import { useMemo } from "preact/hooks";

interface Resource {
    total: number;
    used: number
}

interface ResourceLabel {
    label: number;
    usage: number;
}

export interface ResourceInfoResponse {
    memory: Resource;
    cpuUsage: number;
    cpuCores: ResourceLabel[];
}

export async function getResourceInfo() {
    return api.get<ResourceInfoResponse>("/resource-info");
}

export function useResourceInfo() {
    const queryInfo = useQuery({
        queryFn: getResourceInfo,
        queryKey: ["resourceinfo"],
        refetchInterval: Duration.Second * 5
    });

    const resourceInfo = useMemo(() => {
        return queryInfo.data?.data;
    }, [queryInfo.data]);

    const cores = useMemo(() => {
        return resourceInfo?.cpuCores ?? [];
    }, [resourceInfo]);

    return {
        ...queryInfo,
        resourceInfo,
        cores
    };
}