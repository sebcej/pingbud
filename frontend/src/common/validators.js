import { isValidCron } from 'cron-validator'

export function validCron (value) {
    return isValidCron(value, { seconds: true }) || "Invalid cron notation";
}

export function requiredField(value) {
    return value !== "" || "Field is required"
}

export function validIp(value) {
    return /^\d+\.\d+\.\d+\.\d+$/.test(value) || "Invalid ip"
}