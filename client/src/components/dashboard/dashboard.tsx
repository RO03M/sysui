import { CpuProgressBar } from "./cpu-progress-bar";
import { useResourceInfo } from "../../api/get-resource-info";
import { CpuCores } from "../cpu-cores";

export function Dashboard() {
    const { resourceInfo } = useResourceInfo();

    if (!resourceInfo) {
        return <span>No data available</span>
    }

    return (
        <div>
            <CpuProgressBar label={"cpu"} percentage={resourceInfo.cpuUsage}/>
            <CpuCores />
        </div>
    )
}