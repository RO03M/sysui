import { useMemo } from "preact/hooks";
import { ProgressBar } from "../progress-bar";
import styles from "./styles/cpu-progress-bar.module.css";

interface CpuProgressBarProps {
    label: string;
    percentage: number;
}

export function CpuProgressBar(props: CpuProgressBarProps) {
    const usageLabel = useMemo(() => {
        return props.percentage.toFixed(2);
    }, [props.percentage]);

    return (
        <div className={styles.container}>
            <span style={{ minWidth: 100 }}>{props.label}</span>
            <ProgressBar
                percentage={props.percentage}
            />
            <span style={{ minWidth: 70 }}>{usageLabel}%</span>
        </div>
    )
}