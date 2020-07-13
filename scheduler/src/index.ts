import app from './app';

const port = process.env.PORT || 3000;

app.listen(port, (err) => {
    if (err) {
        return console.log(err)
    }

    return console.log(`IO-SDK REST Api Scheduler is running on ${port}`)
})
