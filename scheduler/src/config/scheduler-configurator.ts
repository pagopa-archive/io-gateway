import fs from 'fs';
import JobData from "../model/job-data";

/**
 * Stores and Retrieves scheduling configurations from a JSON File.
 *
 * Path to the config file can be parametrized using the environment variable IO_SDK_SCHEDULER_CONFIG.
 *
 * If not passed it fallback to the default value '${HOME}/io-sdk-scheduler.json
 */
class SchedulerConfigurator {

    /* fallback default config filename if none is provided via */
    private DEFAULT_CONFIG_FILENAME : string = process.env.HOME+'/io-sdk-scheduler-config.json';
    private configFilename: string;

    public constructor(){

        if(process.env.IO_SDK_SCHEDULER_CONFIG) {
            this.configFilename = process.env.IO_SDK_SCHEDULER_CONFIG
        }   else {
            this.configFilename = this.DEFAULT_CONFIG_FILENAME;
        }

        console.log('Configuration store set to file ',this.configFilename);
    }

    /**
     * Sync call to read the file.
     */
    private parsePersitedJobs(): any {
            const data = fs.readFileSync(this.configFilename);
            return JSON.parse(data.toString());
    }

    public retrievePersistedJobs(): Array<JobData> {
        if(fs.existsSync(this.configFilename)){
            console.log('found configuration file ',this.configFilename);
            return this.parsePersitedJobs();
        } else {
            console.log('configuration file does not exists returning empty list of jobs');
            return [];
        }
    }

    /**
     * Sync call to persist the JSON file.
     * @param jobs
     */
    public persistJobs(jobs: Array<JobData>) {
        const jsonContent = JSON.stringify(jobs);
        fs.writeFileSync(this.configFilename,jsonContent);
        console.log('JSON Configuration saved to ',this.configFilename);
    }

    public getConfigFilename(): string {
        return this.configFilename;
    }

}

export default SchedulerConfigurator;
