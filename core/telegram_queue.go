package core

import (
    "sync"
)

// TelegramJob represents a queued Telegram notification
type TelegramJob struct {
    Session  TSession
    ChatID   string
    BotToken string
}

// TelegramQueue manages async Telegram notifications
type TelegramQueue struct {
    jobs    chan TelegramJob
    wg      sync.WaitGroup
    started bool
}

var (
    defaultQueue *TelegramQueue
    queueOnce    sync.Once
)

// GetTelegramQueue returns the singleton Telegram queue
func GetTelegramQueue() *TelegramQueue {
    queueOnce.Do(func() {
        defaultQueue = &TelegramQueue{
            jobs: make(chan TelegramJob, 100), // buffer up to 100 jobs
        }
    })
    return defaultQueue
}

// Start begins processing the notification queue
func (q *TelegramQueue) Start() {
    if q.started {
        return
    }
    q.started = true
    
    q.wg.Add(1)
    go func() {
        defer q.wg.Done()
        for job := range q.jobs {
            // Process each job synchronously within the goroutine
            Notify(job.Session, job.ChatID, job.BotToken)
        }
    }()
    
    log.Debug("telegram: notification queue started")
}

// Stop gracefully shuts down the queue
func (q *TelegramQueue) Stop() {
    if q.started {
        close(q.jobs)
        q.wg.Wait()
        q.started = false
        log.Debug("telegram: notification queue stopped")
    }
}

// Enqueue adds a notification job to the queue
func (q *TelegramQueue) Enqueue(session TSession, chatID, botToken string) {
    q.jobs <- TelegramJob{
        Session:  session,
        ChatID:   chatID,
        BotToken: botToken,
    }
}
