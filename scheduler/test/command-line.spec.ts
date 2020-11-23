import chai from 'chai';
import CommandLine from "../src/command/command-line";

describe("Command Line",function(){
    describe("run", function(){
        it('should execute the specified command and store the stdout', async function () {
               const commandLine = new CommandLine("echo this is a test message");
               await commandLine.run();
               chai.expect("this is a test message\n").to.equal(commandLine.getCommandResult());
        });
    })
});
