import process from 'child_process';
import util from 'util';
import Runnable from './runnable';

/**
 * A command line executor
 */
class CommandLine implements Runnable {

    command: string;
    stdout: string;
    stderr: string;

    constructor (command: string) {
        this.command = command;
    }

    /**
     * Kept for reference.
     */
    execute(){
        return new Promise((resolve, reject) => {
           process.exec(this.command,(error, stdout, stderr) => {
              if(error){
                  reject();
                  return;
              }

              if(stderr) {
                  reject(stderr);
                  return;
              }
              resolve(stdout);
           });
        });
    }

    public getCommandResult(): string {
        return this.stdout;
    }

    /**
     * Execute the provided command.
     */
    public async run(){
        console.log('Executing command line ',this.command);
        await this.execute().then((stdout: string)=>{
            this.stdout = stdout;
            console.log('Completed execution of command ',this.command);
        }).catch((error: string) =>{
            console.log('Failed execution of command ',this.command);
            this.stderr = error;
        });
    }
}

export default CommandLine;
