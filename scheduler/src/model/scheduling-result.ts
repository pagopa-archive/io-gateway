import JobData from "./job-data";

class SchedulingResult {
    public job: JobData;
    public status: string;
    public error: boolean;
    public errorMessage: string;
}

export default SchedulingResult;
